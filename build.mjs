import { build, context } from 'esbuild';
import { copyFile, readdir, readFile, mkdir, writeFile } from 'node:fs/promises';

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

// Copy Lucide SVG icons into dist/icons/ so Go can embed them.
async function copyIcons() {
  const src = 'node_modules/lucide-static/icons';
  const dest = 'dist/icons';
  await mkdir(dest, { recursive: true });
  const files = await readdir(src);
  let count = 0;
  for (const f of files) {
    if (!f.endsWith('.svg')) continue;
    const svg = await readFile(`${src}/${f}`, 'utf8');
    // Strip the license comment to keep embeds small.
    const clean = svg.replace(/<!--[\s\S]*?-->\n?/, '').trim();
    await writeFile(`${dest}/${f}`, clean);
    count++;
  }
  // The per-file license comments are stripped above, so ship the license once.
  await copyFile('node_modules/lucide-static/LICENSE', `${dest}/LICENSE`);
  console.log(`  dist/icons/   ${count} SVGs`);
}

if (process.argv.includes('--watch')) {
  const jsCtx = await context(jsOpts);
  const cssCtx = await context(cssOpts);
  await jsCtx.watch();
  await cssCtx.watch();
  console.log('watching src/');
} else {
  await build(jsOpts);
  await build(cssOpts);
  await copyIcons();
}
