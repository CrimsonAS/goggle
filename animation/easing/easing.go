/*
 * Copyright (C) 2017 Crimson AS <info@crimson.no>
 * Copyright (C) 2001 Robert Penner
 * All rights reserved.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions are met:
 *
 *  * Redistributions of source code must retain the above copyright notice,
 *    this list of conditions and the following disclaimer.
 *  * Redistributions in binary form must reproduce the above copyright notice,
 *    this list of conditions and the following disclaimer in the documentation
 *    and/or other materials provided with the distribution.
 *  * Neither the name of the author nor the names of contributors may be used
 *    to endorse or promote products derived from this software without
 *    specific prior written permission.
 *
 * THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
 * AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
 * IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
 * ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT OWNER OR CONTRIBUTORS BE
 * LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
 * CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF
 * SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS
 * INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN
 * CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
 * ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
 * POSSIBILITY OF SUCH DAMAGE.
 */

package easing

import "math"

const pi2 = math.Pi / 2

/**
 * Easing equation function for a simple linear tweening, with no easing.
 *
 * @param t		Current time (in frames or seconds).
 * @return		The correct value.
 */
//func None(t float64) float64
//{
//    return t;
//}

/**
 * Easing equation function for a quadratic (t^2) easing in: accelerating from zero velocity.
 *
 * @param t		Current time (in frames or seconds).
 * @return		The correct value.
 */
func InQuad(t float64) float64 {
	return t * t
}

/**
* Easing equation function for a quadratic (t^2) easing out: decelerating to zero velocity.
*
* @param t		Current time (in frames or seconds).
* @return		The correct value.
 */
func OutQuad(t float64) float64 {
	return -t * (t - 2)
}

/**
 * Easing equation function for a quadratic (t^2) easing in/out: acceleration until halfway, then deceleration.
 *
 * @param t		Current time (in frames or seconds).
 * @return		The correct value.
 */
func InOutQuad(t float64) float64 {
	t *= 2.0
	if t < 1 {
		return t * t / 2
	} else {
		t -= 1
		return -0.5 * (t*(t-2) - 1)
	}
}

/**
 * Easing equation function for a quadratic (t^2) easing out/in: deceleration until halfway, then acceleration.
 *
 * @param t		Current time (in frames or seconds).
 * @return		The correct value.
 */
func OutInQuad(t float64) float64 {
	if t < 0.5 {
		return OutQuad(t*2) / 2
	}
	return InQuad((2*t)-1)/2 + 0.5
}

/**
 * Easing equation function for a cubic (t^3) easing in: accelerating from zero velocity.
 *
 * @param t		Current time (in frames or seconds).
 * @return		The correct value.
 */
func InCubic(t float64) float64 {
	return t * t * t
}

/**
 * Easing equation function for a cubic (t^3) easing out: decelerating to zero velocity.
 *
 * @param t		Current time (in frames or seconds).
 * @return		The correct value.
 */
func OutCubic(t float64) float64 {
	t -= 1.0
	return t*t*t + 1
}

/**
 * Easing equation function for a cubic (t^3) easing in/out: acceleration until halfway, then deceleration.
 *
 * @param t		Current time (in frames or seconds).
 * @return		The correct value.
 */
func InOutCubic(t float64) float64 {
	t *= 2.0
	if t < 1 {
		return 0.5 * t * t * t
	} else {
		t -= 2.0
		return 0.5 * (t*t*t + 2)
	}
}

/**
 * Easing equation function for a cubic (t^3) easing out/in: deceleration until halfway, then acceleration.
 *
 * @param t		Current time (in frames or seconds).
 * @return		The correct value.
 */
func OutInCubic(t float64) float64 {
	if t < 0.5 {
		return OutCubic(2*t) / 2
	}
	return InCubic(2*t-1)/2 + 0.5
}

/**
 * Easing equation function for a quartic (t^4) easing in: accelerating from zero velocity.
 *
 * @param t		Current time (in frames or seconds).
 * @return		The correct value.
 */
func InQuart(t float64) float64 {
	return t * t * t * t
}

/**
 * Easing equation function for a quartic (t^4) easing out: decelerating to zero velocity.
 *
 * @param t		Current time (in frames or seconds).
 * @return		The correct value.
 */
func OutQuart(t float64) float64 {
	t -= 1.0
	return -(t*t*t*t - 1)
}

/**
 * Easing equation function for a quartic (t^4) easing in/out: acceleration until halfway, then deceleration.
 *
 * @param t		Current time (in frames or seconds).
 * @return		The correct value.
 */
func InOutQuart(t float64) float64 {
	t *= 2
	if t < 1 {
		return 0.5 * t * t * t * t
	} else {
		t -= 2.0
		return -0.5 * (t*t*t*t - 2)
	}
}

/**
 * Easing equation function for a quartic (t^4) easing out/in: deceleration until halfway, then acceleration.
 *
 * @param t		Current time (in frames or seconds).
 * @return		The correct value.
 */
func OutInQuart(t float64) float64 {
	if t < 0.5 {
		return OutQuart(2*t) / 2
	}
	return InQuart(2*t-1)/2 + 0.5
}

/**
 * Easing equation function for a quintic (t^5) easing in: accelerating from zero velocity.
 *
 * @param t		Current time (in frames or seconds).
 * @return		The correct value.
 */
func InQuint(t float64) float64 {
	return t * t * t * t * t
}

/**
 * Easing equation function for a quintic (t^5) easing out: decelerating to zero velocity.
 *
 * @param t		Current time (in frames or seconds).
 * @return		The correct value.
 */
func OutQuint(t float64) float64 {
	t -= 1.0
	return t*t*t*t*t + 1
}

/**
 * Easing equation function for a quintic (t^5) easing in/out: acceleration until halfway, then deceleration.
 *
 * @param t		Current time (in frames or seconds).
 * @return		The correct value.
 */
func InOutQuint(t float64) float64 {
	t *= 2.0
	if t < 1 {
		return 0.5 * t * t * t * t * t
	} else {
		t -= 2.0
		return 0.5 * (t*t*t*t*t + 2)
	}
}

/**
 * Easing equation function for a quintic (t^5) easing out/in: deceleration until halfway, then acceleration.
 *
 * @param t		Current time (in frames or seconds).
 * @return		The correct value.
 */
func OutInQuint(t float64) float64 {
	if t < 0.5 {
		return OutQuint(2*t) / 2
	}
	return InQuint(2*t-1)/2 + 0.5
}

/**
 * Easing equation function for a sinusoidal (sin(t)) easing in: accelerating from zero velocity.
 *
 * @param t		Current time (in frames or seconds).
 * @return		The correct value.
 */
func InSine(t float64) float64 {
	if t == 1.0 {
		return 1.0
	}
	return -math.Cos(t*pi2) + 1.0
}

/**
 * Easing equation function for a sinusoidal (sin(t)) easing out: decelerating to zero velocity.
 *
 * @param t		Current time (in frames or seconds).
 * @return		The correct value.
 */
func OutSine(t float64) float64 {
	return math.Sin(t * pi2)
}

/**
 * Easing equation function for a sinusoidal (sin(t)) easing in/out: acceleration until halfway, then deceleration.
 *
 * @param t		Current time (in frames or seconds).
 * @return		The correct value.
 */
func InOutSine(t float64) float64 {
	return -0.5 * (math.Cos(math.Pi*t) - 1)
}

/**
 * Easing equation function for a sinusoidal (sin(t)) easing out/in: deceleration until halfway, then acceleration.
 *
 * @param t		Current time (in frames or seconds).
 * @return		The correct value.
 */
func OutInSine(t float64) float64 {
	if t < 0.5 {
		return OutSine(2*t) / 2
	}
	return InSine(2*t-1)/2 + 0.5
}

/**
 * Easing equation function for an exponential (2^t) easing in: accelerating from zero velocity.
 *
 * @param t		Current time (in frames or seconds).
 * @return		The correct value.
 */
func InExpo(t float64) float64 {
	if t == 0 || t == 1.0 {
		return t
	}
	return math.Pow(2.0, 10*(t-1)) - 0.001
}

/**
 * Easing equation function for an exponential (2^t) easing out: decelerating to zero velocity.
 *
 * @param t		Current time (in frames or seconds).
 * @return		The correct value.
 */
func OutExpo(t float64) float64 {
	if t == 1.0 {
		return 1.0
	}
	return 1.001 * (-math.Pow(2.0, -10*t) + 1)
}

/**
 * Easing equation function for an exponential (2^t) easing in/out: acceleration until halfway, then deceleration.
 *
 * @param t		Current time (in frames or seconds).
 * @return		The correct value.
 */
func InOutExpo(t float64) float64 {
	if t == 0.0 {
		return 0.0
	}
	if t == 1.0 {
		return 1.0
	}
	t *= 2.0
	if t < 1 {
		return 0.5*math.Pow(2.0, 10*(t-1)) - 0.0005
	}
	return 0.5 * 1.0005 * (-math.Pow(2.0, -10*(t-1)) + 2)
}

/**
 * Easing equation function for an exponential (2^t) easing out/in: deceleration until halfway, then acceleration.
 *
 * @param t		Current time (in frames or seconds).
 * @return		The correct value.
 */
func OutInExpo(t float64) float64 {
	if t < 0.5 {
		return OutExpo(2*t) / 2
	}
	return InExpo(2*t-1)/2 + 0.5
}

/**
 * Easing equation function for a circular (sqrt(1-t^2)) easing in: accelerating from zero velocity.
 *
 * @param t		Current time (in frames or seconds).
 * @return		The correct value.
 */
func InCirc(t float64) float64 {
	return -(math.Sqrt(1-t*t) - 1)
}

/**
 * Easing equation function for a circular (sqrt(1-t^2)) easing out: decelerating to zero velocity.
 *
 * @param t		Current time (in frames or seconds).
 * @return		The correct value.
 */
func OutCirc(t float64) float64 {
	t -= 1.0
	return math.Sqrt(1 - t*t)
}

/**
 * Easing equation function for a circular (sqrt(1-t^2)) easing in/out: acceleration until halfway, then deceleration.
 *
 * @param t		Current time (in frames or seconds).
 * @return		The correct value.
 */
func InOutCirc(t float64) float64 {
	t *= 2.0
	if t < 1 {
		return -0.5 * (math.Sqrt(1-t*t) - 1)
	} else {
		t -= 2.0
		return 0.5 * (math.Sqrt(1-t*t) + 1)
	}
}

/**
 * Easing equation function for a circular (sqrt(1-t^2)) easing out/in: deceleration until halfway, then acceleration.
 *
 * @param t		Current time (in frames or seconds).
 * @return		The correct value.
 */
func OutInCirc(t float64) float64 {
	if t < 0.5 {
		return OutCirc(2*t) / 2
	}
	return InCirc(2*t-1)/2 + 0.5
}

func easeInElastic_helper(t, b, c, d, a, p float64) float64 {
	if t == 0 {
		return b
	}
	t_adj := t / d
	if t_adj == 1 {
		return b + c
	}

	s := 0.0
	if a < math.Abs(c) {
		a = c
		s = p / 4.0
	} else {
		s = p / (2 * math.Pi) * math.Asin(c/a)
	}

	t_adj -= 1.0
	return -(a * math.Pow(2.0, 10*t_adj) * math.Sin((t_adj*d-s)*(2*math.Pi)/p)) + b
}

/**
 * Easing equation function for an elastic (exponentially decaying sine wave) easing in: accelerating from zero velocity.
 *
 * @param t		Current time (in frames or seconds).
 * @param a		Amplitude.
 * @param p		Period.
 * @return		The correct value.
 */
func InElastic(t, a, p float64) float64 {
	return easeInElastic_helper(t, 0, 1, 1, a, p)
}

func easeOutElastic_helper(t, _ /*b*/, c, _ /*d*/, a, p float64) float64 {
	if t == 0 {
		return 0
	}
	if t == 1 {
		return c
	}

	s := 0.0
	if a < c {
		a = c
		s = p / 4.0
	} else {
		s = p / (2 * math.Pi) * math.Asin(c/a)
	}

	return (a*math.Pow(2.0, -10*t)*math.Sin((t-s)*(2*math.Pi)/p) + c)
}

/**
 * Easing equation function for an elastic (exponentially decaying sine wave) easing out: decelerating to zero velocity.
 *
 * @param t		Current time (in frames or seconds).
 * @param a		Amplitude.
 * @param p		Period.
 * @return		The correct value.
 */
func OutElastic(t, a, p float64) float64 {
	return easeOutElastic_helper(t, 0, 1, 1, a, p)
}

/**
 * Easing equation function for an elastic (exponentially decaying sine wave) easing in/out: acceleration until halfway, then deceleration.
 *
 * @param t		Current time (in frames or seconds).
 * @param a		Amplitude.
 * @param p		Period.
 * @return		The correct value.
 */
func InOutElastic(t, a, p float64) float64 {
	if t == 0 {
		return 0.0
	}
	t *= 2.0
	if t == 2 {
		return 1.0
	}

	s := 0.0
	if a < 1.0 {
		a = 1.0
		s = p / 4.0
	} else {
		s = p / (2 * math.Pi) * math.Asin(1.0/a)
	}

	if t < 1 {
		return -.5 * (a * math.Pow(2.0, 10*(t-1)) * math.Sin((t-1-s)*(2*math.Pi)/p))
	}
	return a*math.Pow(2.0, -10*(t-1))*math.Sin((t-1-s)*(2*math.Pi)/p)*.5 + 1.0
}

/**
 * Easing equation function for an elastic (exponentially decaying sine wave) easing out/in: deceleration until halfway, then acceleration.
 *
 * @param t		Current time (in frames or seconds).
 * @param a		Amplitude.
 * @param p		Period.
 * @return		The correct value.
 */
func OutInElastic(t, a, p float64) float64 {
	if t < 0.5 {
		return easeOutElastic_helper(t*2, 0, 0.5, 1.0, a, p)
	}
	return easeInElastic_helper(2*t-1.0, 0.5, 0.5, 1.0, a, p)
}

/**
 * Easing equation function for a back (overshooting cubic easing: (s+1)*t^3 - s*t^2) easing in: accelerating from zero velocity.
 *
 * @param t		Current time (in frames or seconds).
 * @param s		Overshoot ammount: higher s means greater overshoot (0 produces cubic easing with no overshoot, and the default value of 1.70158 produces an overshoot of 10 percent).
 * @return		The correct value.
 */
func InBack(t, s float64) float64 {
	return t * t * ((s+1)*t - s)
}

/**
 * Easing equation function for a back (overshooting cubic easing: (s+1)*t^3 - s*t^2) easing out: decelerating to zero velocity.
 *
 * @param t		Current time (in frames or seconds).
 * @param s		Overshoot ammount: higher s means greater overshoot (0 produces cubic easing with no overshoot, and the default value of 1.70158 produces an overshoot of 10 percent).
 * @return		The correct value.
 */
func OutBack(t, s float64) float64 {
	t -= 1.0
	return t*t*((s+1)*t+s) + 1
}

/**
 * Easing equation function for a back (overshooting cubic easing: (s+1)*t^3 - s*t^2) easing in/out: acceleration until halfway, then deceleration.
 *
 * @param t		Current time (in frames or seconds).
 * @param s		Overshoot ammount: higher s means greater overshoot (0 produces cubic easing with no overshoot, and the default value of 1.70158 produces an overshoot of 10 percent).
 * @return		The correct value.
 */
func InOutBack(t, s float64) float64 {
	t *= 2.0
	if t < 1 {
		s *= 1.525
		return 0.5 * (t * t * ((s+1)*t - s))
	} else {
		t -= 2
		s *= 1.525
		return 0.5 * (t*t*((s+1)*t+s) + 2)
	}
}

/**
 * Easing equation function for a back (overshooting cubic easing: (s+1)*t^3 - s*t^2) easing out/in: deceleration until halfway, then acceleration.
 *
 * @param t		Current time (in frames or seconds).
 * @param s		Overshoot ammount: higher s means greater overshoot (0 produces cubic easing with no overshoot, and the default value of 1.70158 produces an overshoot of 10 percent).
 * @return		The correct value.
 */
func OutInBack(t, s float64) float64 {
	if t < 0.5 {
		return OutBack(2*t, s) / 2
	}
	return InBack(2*t-1, s)/2 + 0.5
}

func easeOutBounce_helper(t, c, a float64) float64 {
	if t == 1.0 {
		return c
	} else if t < (4 / 11.0) {
		return c * (7.5625 * t * t)
	} else if t < (8 / 11.0) {
		t -= (6 / 11.0)
		return -a*(1.-(7.5625*t*t+.75)) + c
	} else if t < (10 / 11.0) {
		t -= (9 / 11.0)
		return -a*(1.-(7.5625*t*t+.9375)) + c
	} else {
		t -= (21 / 22.0)
		return -a*(1.-(7.5625*t*t+.984375)) + c
	}
}

/**
 * Easing equation function for a bounce (exponentially decaying parabolic bounce) easing out: decelerating to zero velocity.
 *
 * @param t		Current time (in frames or seconds).
 * @param a		Amplitude.
 * @return		The correct value.
 */
func OutBounce(t, a float64) float64 {
	return easeOutBounce_helper(t, 1, a)
}

/**
 * Easing equation function for a bounce (exponentially decaying parabolic bounce) easing in: accelerating from zero velocity.
 *
 * @param t		Current time (in frames or seconds).
 * @param a		Amplitude.
 * @return		The correct value.
 */
func InBounce(t, a float64) float64 {
	return 1.0 - easeOutBounce_helper(1.0-t, 1.0, a)
}

/**
 * Easing equation function for a bounce (exponentially decaying parabolic bounce) easing in/out: acceleration until halfway, then deceleration.
 *
 * @param t		Current time (in frames or seconds).
 * @param a		Amplitude.
 * @return		The correct value.
 */
func InOutBounce(t, a float64) float64 {
	if t < 0.5 {
		return InBounce(2*t, a) / 2
	} else if t == 1.0 {
		return 1.0
	}
	return OutBounce(2*t-1, a)/2 + 0.5
}

/**
 * Easing equation function for a bounce (exponentially decaying parabolic bounce) easing out/in: deceleration until halfway, then acceleration.
 *
 * @param t		Current time (in frames or seconds).
 * @param a		Amplitude.
 * @return		The correct value.
 */
func OutInBounce(t, a float64) float64 {
	if t < 0.5 {
		return easeOutBounce_helper(t*2, 0.5, a)
	}
	return 1.0 - easeOutBounce_helper(2.0-2*t, 0.5, a)
}

func sinProgress(value float64) float64 {
	return math.Sin((value*math.Pi)-pi2)/2 + 0.5
}

func smoothBeginEndMixFactor(value float64) float64 {
	return math.Min(math.Max(1-value*2+0.3, 0.0), 1.0)
}

// SmoothBegin blends Smooth and Linear Interpolation.
// Progress 0 - 0.3      : Smooth only
// Progress 0.3 - ~ 0.5  : Mix of Smooth and Linear
// Progress ~ 0.5  - 1   : Linear only

/**
 * Easing function that starts growing slowly, then incrEases in speed. At the end of the curve the speed will be constant.
 */
func InCurve(t float64) float64 {
	sinProgress := sinProgress(t)
	mix := smoothBeginEndMixFactor(t)
	return sinProgress*mix + t*(1-mix)
}

/**
 * Easing function that starts growing steadily, then ends slowly. The speed will be constant at the beginning of the curve.
 */
func OutCurve(t float64) float64 {
	sinProgress := sinProgress(t)
	mix := smoothBeginEndMixFactor(1 - t)
	return sinProgress*mix + t*(1-mix)
}

/**
 * Easing function where the value grows sinusoidally. Note that the calculated  end value will be 0 rather than 1.
 */
func SineCurve(t float64) float64 {
	return (math.Sin((t*math.Pi*2)-pi2) + 1) / 2
}

/**
 * Easing function where the value grows cosinusoidally. Note that the calculated start value will be 0.5 and the end value will be 0.5
 * contrary to the usual 0 to 1 easing curve.
 */
func CosineCurve(t float64) float64 {
	return (math.Cos((t*math.Pi*2)-pi2) + 1) / 2
}
