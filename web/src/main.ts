import "./style.css";

const app = document.querySelector<HTMLDivElement>("#app")!;

app.innerHTML = `
  <main>
    <h1>Retro Games</h1>

    <canvas id="game" width="281" height="421"></canvas>
  </main>
`;

const go = new Go();

const canvas = document.getElementById("game") as HTMLCanvasElement;
const context = canvas.getContext("2d")!;
context.imageSmoothingEnabled = false;

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
