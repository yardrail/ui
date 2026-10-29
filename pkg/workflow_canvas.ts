import rough from 'roughjs';

export class YrWorkflowCanvas extends HTMLElement {
  private svg: SVGSVGElement | null = null;
  private observer: ResizeObserver | null = null;

  connectedCallback() {
    this.svg = document.createElementNS('http://www.w3.org/2000/svg', 'svg');
    this.svg.classList.add('yr-workflow-canvas__grid');
    this.prepend(this.svg);
    this.drawGrid();

    this.observer = new ResizeObserver(() => this.drawGrid());
    this.observer.observe(this);
  }

  disconnectedCallback() {
    this.observer?.disconnect();
  }

  private drawGrid() {
    const svg = this.svg;
    if (!svg) return;

    const w = this.clientWidth;
    const h = this.clientHeight;
    const spacing = 100;

    svg.setAttribute('width', String(w));
    svg.setAttribute('height', String(h));
    svg.setAttribute('viewBox', `0 0 ${w} ${h}`);
    svg.innerHTML = '';

    const rc = rough.svg(svg);
    const opts = {
      stroke: '#d0cdc6',
      strokeWidth: 1,
      roughness: 1.5,
      bowing: 1,
      disableMultiStroke: true,
    };

    for (let x = spacing; x < w; x += spacing) {
      svg.appendChild(rc.line(x, 0, x, h, opts));
    }

    for (let y = spacing; y < h; y += spacing) {
      svg.appendChild(rc.line(0, y, w, y, opts));
    }
  }
}

customElements.define('yr-workflow-canvas', YrWorkflowCanvas);
