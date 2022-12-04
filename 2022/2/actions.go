package main

type Action interface {
	Outcome(action Action) (outcome Outcome)
	Points() int
	RequiredActionFor(outcome Outcome) Action
}

type Rock struct {
}

func (r Rock) RequiredActionFor(outcome Outcome) Action {
	if _, ok := outcome.(Win); ok {
		return Paper{}
	} else if _, ok := outcome.(Loss); ok {
		return Scissors{}
	}
	return Rock{}
}

func (r Rock) Outcome(action Action) Outcome {
	if _, ok := action.(Rock); ok {
		return Draw{}
	} else if _, ok := action.(Paper); ok {
		return Loss{}
	}
	return Win{}
}

func (r Rock) Points() int {
	return 1
}

type Paper struct {
}

func (p Paper) RequiredActionFor(outcome Outcome) Action {
	if _, ok := outcome.(Win); ok {
		return Scissors{}
	} else if _, ok := outcome.(Loss); ok {
		return Rock{}
	}
	return Paper{}
}

func (p Paper) Outcome(action Action) (outcome Outcome) {
	if _, ok := action.(Paper); ok {
		return Draw{}
	} else if _, ok := action.(Scissors); ok {
		return Loss{}
	}
	return Win{}
}

func (p Paper) Points() int {
	return 2
}

type Scissors struct {
}

func (s Scissors) RequiredActionFor(outcome Outcome) Action {
	if _, ok := outcome.(Win); ok {
		return Rock{}
	} else if _, ok := outcome.(Loss); ok {
		return Paper{}
	}
	return Scissors{}
}

func (s Scissors) Outcome(action Action) (outcome Outcome) {
	if _, ok := action.(Scissors); ok {
		return Draw{}
	} else if _, ok := action.(Rock); ok {
		return Loss{}
	}
	return Win{}
}

func (s Scissors) Points() int {
	return 3
}
