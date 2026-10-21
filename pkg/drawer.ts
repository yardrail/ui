export class YrDrawer extends HTMLElement {
  static observedAttributes = ['open'];

  private onKeyDown = (e: KeyboardEvent) => {
    if (e.key === 'Escape' && this.hasAttribute('open')) {
      this.removeAttribute('open');
    }
  };

  connectedCallback() {
    document.addEventListener('keydown', this.onKeyDown);
  }

  disconnectedCallback() {
    document.removeEventListener('keydown', this.onKeyDown);
  }

  attributeChangedCallback() {
    // CSS handles all visual changes via yr-drawer[open]
  }
}

export class YrDrawerHeader extends HTMLElement {}
export class YrDrawerBody extends HTMLElement {}

export class YrDrawerClose extends HTMLElement {
  private onClick = () => {
    this.closest('yr-drawer')?.removeAttribute('open');
  };

  connectedCallback() {
    this.setAttribute('role', 'button');
    this.setAttribute('tabindex', '0');
    this.setAttribute('aria-label', 'Close drawer');
    this.addEventListener('click', this.onClick);
    this.addEventListener('keydown', this.onKeyDown);
  }

  disconnectedCallback() {
    this.removeEventListener('click', this.onClick);
    this.removeEventListener('keydown', this.onKeyDown);
  }

  private onKeyDown = (e: KeyboardEvent) => {
    if (e.key === 'Enter' || e.key === ' ') {
      e.preventDefault();
      this.onClick();
    }
  };
}

customElements.define('yr-drawer', YrDrawer);
customElements.define('yr-drawer-header', YrDrawerHeader);
customElements.define('yr-drawer-body', YrDrawerBody);
customElements.define('yr-drawer-close', YrDrawerClose);
