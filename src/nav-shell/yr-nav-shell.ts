export class YrNavShell extends HTMLElement {
  connectedCallback() {
    this.setAttribute('role', 'navigation');
    this.restructure();
  }

  private restructure() {
    if (this.querySelector('yr-nav-shell-right')) return;

    const left = this.querySelector('[slot="left"]');
    const center = this.querySelector('[slot="center"]');

    const leftWrapper = document.createElement('yr-nav-shell-left');
    if (left) leftWrapper.appendChild(left);

    const centerWrapper = document.createElement('yr-nav-shell-center');
    if (center) centerWrapper.appendChild(center);

    const rightWrapper = document.createElement('yr-nav-shell-right');
    const remaining = Array.from(this.childNodes).filter(
      (n) => n !== leftWrapper && n !== centerWrapper,
    );
    for (const child of remaining) rightWrapper.appendChild(child);

    this.replaceChildren(leftWrapper, centerWrapper, rightWrapper);
  }
}

export class YrNavShellLeft extends HTMLElement {}
export class YrNavShellCenter extends HTMLElement {}
export class YrNavShellRight extends HTMLElement {}

customElements.define('yr-nav-shell', YrNavShell);
customElements.define('yr-nav-shell-left', YrNavShellLeft);
customElements.define('yr-nav-shell-center', YrNavShellCenter);
customElements.define('yr-nav-shell-right', YrNavShellRight);
