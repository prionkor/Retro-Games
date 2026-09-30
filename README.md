# Retro Games

A small personal project to recreate some of the simple pixel-based games I remember from childhood.

I wrote about the story behind the project and why I started building it on my blog:
[Retro Games - Rebuilding a Piece of my Childhood](https://sisir.me/blog/retro-games-rebuilding-a-piece-of-my-childhood/?utm_source=chatgpt.com)

## About the Project

This is not intended to be a serious game engine or a production-ready game framework. It is mostly an experiment and a way for me to explore game development by building things from the ground up.

The first game is a simple retro racing game inspired by the old handheld "Brick Games" I played as a child.

The project is written in Go and currently runs in the browser through WebAssembly.

## Building It Step by Step

One of the things I want to explore with this project is how software abstractions can emerge naturally during development.

I am deliberately not trying to design the entire engine upfront.

Instead, I start with the simplest thing that works and introduce an abstraction when the existing design starts becoming limiting.

For example, the project initially had a simple framebuffer. As the game evolved, it became clear that different parts of the screen might need different visual characteristics. That led to the idea of `FramePart`, and eventually to `PixelBlock`.

The same approach applies to other parts of the engine. I want the architecture to be shaped by actual requirements rather than by trying to predict every possible requirement before writing the code.

## Architecture

The project currently has a small separation between the game, engine, and platform layers.

```text
Game
  ↓
Engine
  ↓
Platform
```

The game contains the game-specific logic and owns its visual `FramePart`s.

The engine owns the overall `Frame` and provides the common concepts needed by games.

The platform is responsible for presenting the frame on a particular device or environment.

The goal is to keep the game itself unaware of the platform it is running on. The current platform is the browser, using WebAssembly, but the engine is not designed specifically around the browser.

## Visual Representation

The game uses logical pixels rather than drawing directly to the physical display.

A `Frame` can contain multiple `FramePart`s. Each part has its own position, dimensions, and pixel representation.

A `PixelBlock` describes how one logical pixel should actually look when rendered. This allows different parts of the screen to use different pixel sizes or patterns.

This started from a simple 6×6 pixel pattern and gradually evolved as the need for smaller elements, such as score digits, appeared.

## Current State

The project is actively being developed.

The first game currently includes:

* A simple pixel-based racing track
* Player-controlled car
* Enemy cars
* Collision detection
* Game-over handling
* Curtain-style restart animation
* A small game engine
* Browser support through WebAssembly

There is still a lot to build, and some of the architecture will probably change as the project grows.

That is part of the experiment.

## Why I'm Building This

Mostly because it is fun.

I am not a game developer, and I am not trying to turn this into a commercial game. I wanted to recreate something from my childhood while learning how the pieces of a simple game fit together.

I'm also interested in seeing how a small system evolves when abstractions are introduced only when they become useful.

The project is intentionally small, imperfect, and still changing.

I'll keep building it and see where it goes.
