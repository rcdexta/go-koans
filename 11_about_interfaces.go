package go_koans

func aboutInterfaces() {
	bob := new(Human)     // bob is a kind of *human
	rspec := new(Program) // rspec is a kind of *program

	assert(Runner(bob) == bob) // conformed interfaces need not be declared, they are inferred

	assert(bob.milesCompleted == 0)
	assert(rspec.executionCount == 0)

	runTwice(bob)   // bob fits the profile for a 'runner'
	runTwice(rspec) // rspec also fits the profile for a 'runner'

	assert(bob.milesCompleted == 2)   // bob is affected by running in his own unique way (probably fatigue)
	assert(rspec.executionCount == 2) // rspec can run completely differently than bob, thanks to interfaces
}

// abstract interface and function that requires it

type Runner interface {
	run()
}

func runTwice(r Runner) {
	r.run()
	r.run()
}

// concrete type implementing the interface

type Human struct {
	milesCompleted int
}

func (h *Human) run() {
	h.milesCompleted++
}

// another concrete type implementing the interface

type Program struct {
	executionCount int
}

func (p *Program) run() {
	p.executionCount++
}
