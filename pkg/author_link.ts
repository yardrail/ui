// <yr-author-link> — the full structure (avatar, name, handle, verified badge)
// is server-rendered by the templ component. This class exists only to
// register the custom element tag.
export class YrAuthorLink extends HTMLElement {}

customElements.define('yr-author-link', YrAuthorLink);
