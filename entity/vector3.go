// /////////////////////////////////////////////////////////////////////////////
// 3d 坐标

package entity

import (
	"fmt"
	"math"
)

// Coord is the of coordinations entity position (x, y, z)
type Coord float32

// /////////////////////////////////////////////////////////////////////////////
// Vector3 对象

// 实体 3d 坐标
type Vector3 struct {
	X Coord
	Y Coord
	Z Coord
}

// 打印信息
func (this Vector3) String() string {
	return fmt.Sprintf("(%.2f, %.2f, %.2f)", this.X, this.Y, this.Z)
}

// DistanceTo calculates distance between two positions
func (this Vector3) DistanceTo(o Vector3) Coord {
	dx := this.X - o.X
	dy := this.Y - o.Y
	dz := this.Z - o.Z
	return Coord(math.Sqrt(float64(dx*dx + dy*dy + dz*dz)))
}

// Sub calculates Vector3 p - Vector3 o
func (this Vector3) Sub(o Vector3) Vector3 {
	return Vector3{this.X - o.X, this.Y - o.Y, this.Z - o.Z}
}

func (this Vector3) Add(o Vector3) Vector3 {
	return Vector3{this.X + o.X, this.Y + o.Y, this.Z + o.Z}
}

// Mul calculates Vector3 p * m
func (this Vector3) Mul(m Coord) Vector3 {
	return Vector3{this.X * m, this.Y * m, this.Z * m}
}

// DirToYaw convert direction represented by Vector3 to Yaw
func (dir Vector3) DirToYaw() Yaw {
	dir.Normalize()

	yaw := math.Acos(float64(dir.X))
	if dir.Z < 0 {
		yaw = math.Pi*2 - yaw
	}

	yaw = yaw / math.Pi * 180 // convert to angle

	if yaw <= 90 {
		yaw = 90 - yaw
	} else {
		yaw = 90 + (360 - yaw)
	}

	return Yaw(yaw)
}

func (this *Vector3) Normalize() {
	d := Coord(math.Sqrt(float64(this.X*this.X + this.Y + this.Y + this.Z*this.Z)))
	if d == 0 {
		return
	}
	this.X /= d
	this.Y /= d
	this.Z /= d
}

func (this Vector3) Normalized() Vector3 {
	this.Normalize()
	return this
}
