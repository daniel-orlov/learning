package main

type rogue struct{}

func (r *rogue) makeKettlebell() iKettlebell {
	return &rogueKettlebell{kettlebell{
		color:  "purple",
		weight: 20.0,
	}}
}

func (r *rogue) makeBarbell() iBarbell {
	return &rogueBarbell{barbell{
		purpose: "women",
		weight:  15.0,
	}}
}

type uralGirya struct{}

func (u *uralGirya) makeKettlebell() iKettlebell {
	return &uralGiryaKettlebell{kettlebell{
		color:  "red",
		weight: 32.0,
	}}
}

func (u *uralGirya) makeBarbell() iBarbell {
	return &uralGiryaBarbell{barbell{
		purpose: "olympic weightlifting",
		weight:  20.0,
	}}
}
