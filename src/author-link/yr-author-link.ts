// <yr-author-link href="/authors/x" name="Yardrail" handle="yardrail" verified
//   avatar="YR" size="lg"></yr-author-link> — a consistent way to render an
// author's identity: optional avatar (composes yr-avatar), name and/or
// handle, and a verified badge. Omit `href` for a non-interactive rendering
// (e.g. the author's own profile header — add `heading` there so the name
// renders as an <h1>, since it's that page's title); include `href` to
// render as a link (e.g. a byline elsewhere on the site). All content is
// built from attributes as plain text (no innerHTML of caller data),
// following the yr-stat pattern.
const VERIFIED_BADGE_PATH =
  'M12 22s8-4 8-10V5l-8-3-8 3v7c0 6 8 10 8 10z';

export class YrAuthorLink extends HTMLElement {
  static observedAttributes = ['href', 'name', 'handle', 'verified', 'avatar', 'size', 'heading'];

  connectedCallback() {
    this.render();
  }

  attributeChangedCallback() {
    this.render();
  }

  private render() {
    this.replaceChildren();

    const name = this.getAttribute('name');
    const handle = this.getAttribute('handle');
    const verified = this.hasAttribute('verified');
    const avatarInitials = this.getAttribute('avatar');
    const href = this.getAttribute('href');
    const isLarge = this.getAttribute('size') === 'lg';

    const inner = document.createElement(href ? 'a' : 'span');
    inner.className = 'yr-author-link-inner';
    if (href) (inner as HTMLAnchorElement).href = href;

    if (avatarInitials) {
      const avatar = document.createElement('yr-avatar');
      if (isLarge) avatar.setAttribute('size', 'lg');
      avatar.textContent = avatarInitials;
      inner.appendChild(avatar);
    }

    const badge = (): HTMLElement => {
      const el = document.createElement('span');
      el.className = 'yr-author-link-badge';
      el.title = 'Verified author';
      const svgSize = isLarge ? 16 : 12;
      el.innerHTML = `<svg width="${svgSize}" height="${svgSize}" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round"><path d="${VERIFIED_BADGE_PATH}"></path></svg>`;
      return el;
    };

    const nameTag = this.hasAttribute('heading') ? 'h1' : 'strong';

    if (isLarge && name && handle) {
      const text = document.createElement('div');
      text.className = 'yr-author-link-text';

      const nameRow = document.createElement('div');
      nameRow.className = 'yr-author-link-name-row';
      const nameEl = document.createElement(nameTag);
      nameEl.className = 'yr-author-link-name';
      nameEl.textContent = name;
      nameRow.appendChild(nameEl);
      if (verified) nameRow.appendChild(badge());
      text.appendChild(nameRow);

      const handleEl = document.createElement('p');
      handleEl.className = 'yr-author-link-handle';
      handleEl.textContent = `@${handle}`;
      text.appendChild(handleEl);

      inner.appendChild(text);
    } else if (name) {
      const nameEl = document.createElement(nameTag);
      nameEl.className = 'yr-author-link-name';
      nameEl.textContent = name;
      inner.appendChild(nameEl);
      if (verified) inner.appendChild(badge());
    } else if (handle) {
      inner.appendChild(document.createTextNode(`@${handle}`));
    }

    this.appendChild(inner);
  }
}

customElements.define('yr-author-link', YrAuthorLink);
