package g

import (
	"math/rand"
	"testing"

	"github.com/go-gl/mathgl/mgl32"
)

func mat4(m mgl32.Mat4) Mat4 {
	return Mat4{
		m[0], m[1], m[2], m[3],
		m[4], m[5], m[6], m[7],
		m[8], m[9], m[10], m[11],
		m[12], m[13], m[14], m[15],
	}
}

func R3() Vec3 {
	return V3(
		rand.Float32()*100-50,
		rand.Float32()*100-50,
		rand.Float32()*100-50,
	)
}

func TestLookAtVCross(t *testing.T) {
	for i := 0; i < 10; i++ {
		eye, center, up := R3(), R3(), R3()

		got := LookAtV(eye, center, up)
		expect := mat4(mgl32.LookAt(eye.X, eye.Y, eye.Z, center.X, center.Y, center.Z, up.X, up.Y, up.Z))

		if !got.EqAlmost(expect, 0.1) {
			t.Fatal(eye, center, up, "\ngot:", got, "\nexp:", expect)
		}
	}
}
