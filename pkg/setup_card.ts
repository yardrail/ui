// <yr-setup-card> and sub-elements — all structure is server-rendered by
// the templ component. These classes exist only to register the custom
// element tags.
export class YrSetupCard extends HTMLElement {}
export class YrSetupCardProgress extends HTMLElement {}
export class YrSetupCardHeader extends HTMLElement {}
export class YrSetupCardBody extends HTMLElement {}
export class YrSetupCardFooter extends HTMLElement {}

customElements.define('yr-setup-card', YrSetupCard);
customElements.define('yr-setup-card-progress', YrSetupCardProgress);
customElements.define('yr-setup-card-header', YrSetupCardHeader);
customElements.define('yr-setup-card-body', YrSetupCardBody);
customElements.define('yr-setup-card-footer', YrSetupCardFooter);
