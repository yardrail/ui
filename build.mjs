import { build, context } from 'esbuild';

const jsOpts = {
  entryPoints: ['index.ts'],
  bundle: true,
  format: 'esm',
  target: 'es2022',
  outfile: 'pkg/dist/ui.js',
  sourcemap: true,
  logLevel: 'info',
};

const cssOpts = {
  entryPoints: ['styles.css'],
  bundle: true,
  outfile: 'pkg/dist/ui.css',
  sourcemap: true,
  logLevel: 'info',
  external: ['/fonts/*'],
};

if (process.argv.includes('--watch')) {
  const jsCtx = await context(jsOpts);
  const cssCtx = await context(cssOpts);
  await jsCtx.watch();
  await cssCtx.watch();
  console.log('watching...');
} else {
  await build(jsOpts);
  await build(cssOpts);
}
