// <yr-nav-shell> — left/center/right structure is server-rendered by the templ
// component. These classes exist only to register the custom element tags.
export class YrNavShell extends HTMLElement {}
export class YrNavShellLeft extends HTMLElement {}
export class YrNavShellCenter extends HTMLElement {}
export class YrNavShellRight extends HTMLElement {}

customElements.define('yr-nav-shell', YrNavShell);
customElements.define('yr-nav-shell-left', YrNavShellLeft);
customElements.define('yr-nav-shell-center', YrNavShellCenter);
customElements.define('yr-nav-shell-right', YrNavShellRight);
