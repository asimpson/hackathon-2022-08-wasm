let module;
importScripts("wasm_exec.js")
go = new Go();

self.onmessage = async (e) => {
  if (e.data.type === "module") {
    module = e.data.module;
    return;
  }

  if (module) {
    console.log('Message received: ', e.data.path);
    const instance = await WebAssembly.instantiate(module, go.importObject);
    go.run(instance);
    postMessage(parsePath(e.data.path));
  }
}
