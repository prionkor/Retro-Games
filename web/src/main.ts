import "./style.css";

const app = document.querySelector<HTMLDivElement>("#app")!;

app.innerHTML = `
  <main>
    <h1>Retro Games</h1>

    <canvas id="game" width="320" height="480"></canvas>
  </main>
`;

const go = new Go();

declare global {
  var showMessage: (message: string) => void;
}

globalThis.showMessage = (message: string) => {
  console.log(`Game says: ${message}`);
};

const result = await WebAssembly.instantiateStreaming(
  fetch("/game.wasm"),
  go.importObject,
);

go.run(result.instance);
