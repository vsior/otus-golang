package hw06pipelineexecution

type (
	In  = <-chan any
	Out = In
	Bi  = chan any
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	in = func(in In) Out {
		ch := make(Bi)
		go func() {
			defer close(ch)
			select {
			case <-done:
				return
			default:
				processChan(done, in, ch)
			}
		}()
		return ch
	}(in)

	for _, stage := range stages {
		in = ProcessStage(in, done, stage)
	}
	return in
}

func ProcessStage(in In, done In, stage Stage) Out {
	return func(in In) Out {
		ch := make(Bi)
		go func() {
			defer func() {
				close(ch)
				for x := range in {
					_ = x
				}
			}()
			processChan(done, in, ch)
		}()
		return stage(ch)
	}(in)
}

func processChan(done In, in In, out Bi) {
	for {
		select {
		case v, ok := <-in:
			if !ok {
				return
			}
			select {
			case out <- v:
			case <-done:
				return
			}
		case <-done:
			return
		}
	}
}
