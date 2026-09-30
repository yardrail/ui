import { build, context } from 'esbuild';

const jsOpts = {
  entryPoints: ['src/index.ts'],
  bundle: true,
  format: 'esm',
  target: 'es2022',
  outfile: 'dist/ui.js',
  sourcemap: true,
  logLevel: 'info',
};

const cssOpts = {
  entryPoints: ['src/styles.css'],
  bundle: true,
  outfile: 'dist/ui.css',
  sourcemap: true,
  logLevel: 'info',
};

if (process.argv.includes('--watch')) {
  const jsCtx = await context(jsOpts);
  const cssCtx = await context(cssOpts);
  await jsCtx.watch();
  await cssCtx.watch();
  console.log('watching src/');
} else {
  await build(jsOpts);
  await build(cssOpts);
}
