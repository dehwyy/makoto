import wasm_init, {run} from "./pkg/wasm.js";

const runWasm = async () => {
  await wasm_init("./pkg/wasm_bg.wasm")

  run()
}

runWasm()
