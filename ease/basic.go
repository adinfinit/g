package ease

// t = current time

func OutElasticSmall(t, start, end, duration float32) float32 {
	t /= duration
	ts := t * t
	tc := ts * t
	return start + end*(33*tc*ts+-106*ts*ts+126*tc+-67*ts+15*t)
}

func Wobble(t, base, scale, duration float32) float32 {
	t /= duration
	return base + scale*0.0005*(t-1)*(t-1)*t*(44851-224256*t+224256*t*t)
	//return 1.8358224*scale*t*(9.196429-47.13542*t+70.3125*t2-32.38*t2*t) + start
}
