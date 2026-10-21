// <yr-stat> — number + label structure is server-rendered by the templ
// component. This class exists only to register the custom element tag.
export class YrStat extends HTMLElement {}

customElements.define('yr-stat', YrStat);
