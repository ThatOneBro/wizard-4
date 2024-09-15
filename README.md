# wizard-4

A game written in Go for the [WASM-4](https://wasm4.org) fantasy console.

## Goals
- [ ] Experiment with the WASM-4 fantasy console
- [ ] Run the game on a micro-controller like the Raspberry Pi Pico
- [ ] Experiment with expanding multiplayer beyond the 3 designated remote joysticks (attempt to hijack the netplay functionality and create a protocol for reading data from a server via the remote joystick inputs)

## Building

Build the cart by running:

```shell
make
```

Then run it with:

```shell
w4 run build/cart.wasm
```

For more info about setting up WASM-4, see the [quickstart guide](https://wasm4.org/docs/getting-started/setup?code-lang=go#quickstart).

## Links

- [Documentation](https://wasm4.org/docs): Learn more about WASM-4.
- [Snake Tutorial](https://wasm4.org/docs/tutorials/snake/goal): Learn how to build a complete game
  with a step-by-step tutorial.
- [GitHub](https://github.com/aduros/wasm4): Submit an issue or PR. Contributions are welcome!
