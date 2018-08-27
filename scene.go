package main

import (
	"fmt"
	"log"
	"time"

	img "github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type scene struct {
	t    int
	bg   *sdl.Texture
	bird *bird
}

func newScene(r *sdl.Renderer) (*scene, error) {
	bg, err := img.LoadTexture(r, "res/images/back.png")
	if err != nil {
		return nil, fmt.Errorf("Could not drawBackground: %v", err)
	}
	b, err := newBird(r)
	if err != nil {
		return nil, err
	}
	return &scene{bg: bg, bird: b}, nil
}

func (s *scene) run(events <-chan sdl.Event, r *sdl.Renderer) <-chan error {
	errc := make(chan error)

	go func() {
		defer close(errc)
		tick := time.Tick(10 * time.Millisecond)
		for {
			select {
			case e := <-events:
			  if done := s.handleEvent(e); done {
					return
				}
				log.Printf("event: %T", e)
			case <-tick:
				if err := s.paint(r); err != nil {
					errc <- err
				}
			}
		}
	}()
	return errc
}

func (s *scene) handleEvent(event sdl.Event) bool {
	switch event.(type) {
	case *sdl.QuitEvent:
		return true
	default:
		log.Printf("unknown event %T", event)
		return false
	}

}

func (s *scene) paint(r *sdl.Renderer) error {
	s.t++
	r.Clear()
	if err := r.Copy(s.bg, nil, nil); err != nil {
		return fmt.Errorf("Could not copy background: %v", err)
	}
	if err := s.bird.paint(r); err != nil {
		return err
	}

	r.Present()
	return nil
}

func (s *scene) destroy() {
	s.bg.Destroy()
	s.bird.destroy()
}
