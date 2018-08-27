package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

const gravity = 0.25

type bird struct {
	t        int
	textures []*sdl.Texture
	y, speed float64
}

func newBird(r *sdl.Renderer) (*bird, error) {
	var textures []*sdl.Texture
	for i := 1; i <= 4; i++ {
		path := fmt.Sprintf("res/images/frame-%d.png", i)
		texture, err := img.LoadTexture(r, path)
		if err != nil {
			return nil, fmt.Errorf("Could not load texture: %v", err)
		}
		textures = append(textures, texture)
	}
	return &bird{textures: textures, y: 300, speed: 0}, nil
}

func (b *bird) paint(r *sdl.Renderer) error {
	b.t++
	b.y -= b.speed
	if b.y < 0 {
		b.speed = -b.speed
		b.y = 0
	}
	b.speed += gravity
	rect := &sdl.Rect{X: 10, Y: (600 - int32(b.y)) - 42/2, W: 50, H: 43}
	i := b.t / 10 % len(b.textures)
	if err := r.Copy(b.textures[i], nil, rect); err != nil {
		return fmt.Errorf("Could not copy bird: %v", err)
	}
	return nil
}

func (b *bird) destroy() {
	for _, b := range b.textures {
		b.Destroy()
	}
}
