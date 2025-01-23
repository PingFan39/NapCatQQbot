import { encode } from 'silk-wasm';

const audioWorker = async ({ input, sampleRate }) => {
  return await encode(input, sampleRate);
};

export { audioWorker as default };
