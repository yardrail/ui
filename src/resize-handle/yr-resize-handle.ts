export class YrResizeHandle extends HTMLElement {
  static observedAttributes = ['target'];

  private startX = 0;
  private startWidth = 0;

  private get prop(): string {
    return this.getAttribute('target') === 'drawer'
      ? '--yr-drawer-width'
      : '--yr-sidebar-width';
  }

  private get isSidebar(): boolean {
    return this.getAttribute('target') !== 'drawer';
  }

  private resizeTarget(): HTMLElement | null {
    const tag = this.isSidebar ? 'yr-sidebar' : 'yr-drawer';
    const shell = this.closest('body') ?? this.parentElement;
    return shell?.querySelector(tag) as HTMLElement | null;
  }

  private onMouseMove = (e: MouseEvent) => {
    e.preventDefault();
    const delta = this.isSidebar
      ? e.clientX - this.startX
      : this.startX - e.clientX;
    document.documentElement.style.setProperty(
      this.prop,
      `${Math.max(100, this.startWidth + delta)}px`,
    );
  };

  private onMouseUp = () => {
    this.removeAttribute('active');
    document.documentElement.classList.remove('yr-resizing');
    document.body.style.cursor = '';
    document.body.style.userSelect = '';
    document.removeEventListener('mousemove', this.onMouseMove);
    document.removeEventListener('mouseup', this.onMouseUp);
  };

  private onMouseDown = (e: MouseEvent) => {
    e.preventDefault();
    const target = this.resizeTarget();
    if (!target) return;

    this.startX = e.clientX;
    this.startWidth = target.getBoundingClientRect().width;
    this.setAttribute('active', '');
    document.documentElement.classList.add('yr-resizing');
    document.body.style.cursor = 'col-resize';
    document.body.style.userSelect = 'none';
    document.addEventListener('mousemove', this.onMouseMove);
    document.addEventListener('mouseup', this.onMouseUp);
  };

  connectedCallback() {
    this.addEventListener('mousedown', this.onMouseDown);
  }

  disconnectedCallback() {
    this.removeEventListener('mousedown', this.onMouseDown);
    document.removeEventListener('mousemove', this.onMouseMove);
    document.removeEventListener('mouseup', this.onMouseUp);
  }

  attributeChangedCallback() {
    // CSS handles visual changes via yr-resize-handle[target="sidebar"|"drawer"]
  }
}

customElements.define('yr-resize-handle', YrResizeHandle);
