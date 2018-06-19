package g

import "fmt"

type Mat4 struct {
	M00, M01, M02, M03 float32
	M10, M11, M12, M13 float32
	M20, M21, M22, M23 float32
	M30, M31, M32, M33 float32
}

const Mat4Size = 4 * 4 * 4

func (a *Mat4) Ptr() *float32 { return &a.M00 }

func (a Mat4) String() string {
	return fmt.Sprintf("[[%2.2f %2.2f %2.2f %2.2f][%2.2f %2.2f %2.2f %2.2f][%2.2f %2.2f %2.2f %2.2f][%2.2f %2.2f %2.2f %2.2f]]",
		a.M00, a.M01, a.M02, a.M03,
		a.M10, a.M11, a.M12, a.M13,
		a.M20, a.M21, a.M22, a.M23,
		a.M30, a.M31, a.M32, a.M33)
}

func (a Mat4) EqAlmost(b Mat4, eps float32) bool {
	return (Abs(a.M00-b.M00) < eps) && (Abs(a.M01-b.M01) < eps) && (Abs(a.M02-b.M02) < eps) && (Abs(a.M03-b.M03) < eps) &&
		(Abs(a.M10-b.M10) < eps) && (Abs(a.M11-b.M11) < eps) && (Abs(a.M12-b.M12) < eps) && (Abs(a.M13-b.M13) < eps) &&
		(Abs(a.M20-b.M20) < eps) && (Abs(a.M21-b.M21) < eps) && (Abs(a.M22-b.M22) < eps) && (Abs(a.M23-b.M23) < eps) &&
		(Abs(a.M30-b.M30) < eps) && (Abs(a.M31-b.M31) < eps) && (Abs(a.M32-b.M32) < eps) && (Abs(a.M33-b.M33) < eps)
}

func Perspective(fovy, aspect, near, far float32) Mat4 {
	f := 1.0 / Tan(fovy/2.0)
	return Mat4{
		f / aspect, 0, 0, 0,
		0, f, 0, 0,
		0, 0, -(near + far) / (far - near), -1,
		0, 0, -2.0 * far * near / (far - near), 0,
	}
}

func LookAtV(eye, center, up Vec3) Mat4 {
	f := center.Sub(eye).Normalize()
	s := f.Cross(up.Normalize()).Normalize()
	u := s.Cross(f)

	e0, e1, e2 := eye.Dot(s), eye.Dot(u), eye.Dot(f)
	return Mat4{
		s.X, u.X, -f.X, 0,
		s.Y, u.Y, -f.Y, 0,
		s.Z, u.Z, -f.Z, 0,
		-e0, -e1, e2, 1,
	}
}
