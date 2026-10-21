// <yr-navbar> and <yr-navbar-std> — full structure is server-rendered by
// the templ component. These classes exist only to register the custom
// element tags.
export class YrNavbar extends HTMLElement {}
export class YrNavbarStd extends HTMLElement {}

customElements.define('yr-navbar', YrNavbar);
customElements.define('yr-navbar-std', YrNavbarStd);
