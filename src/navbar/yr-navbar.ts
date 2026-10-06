export class YrNavbar extends HTMLElement {
  connectedCallback() {
    this.setAttribute('role', 'navigation');
    this.restructure();
  }

  private restructure() {
    if (this.querySelector('.yr-navbar-left')) return;

    const left = this.querySelector('[slot="left"]');
    const center = this.querySelector('[slot="center"]');

    const leftDiv = document.createElement('div');
    leftDiv.className = 'yr-navbar-left';
    if (left) leftDiv.appendChild(left);

    const centerDiv = document.createElement('div');
    centerDiv.className = 'yr-navbar-center';
    if (center) centerDiv.appendChild(center);

    const rightDiv = document.createElement('div');
    rightDiv.className = 'yr-navbar-right';
    const remaining = Array.from(this.childNodes).filter(
      (n) => n !== leftDiv && n !== centerDiv,
    );
    for (const child of remaining) rightDiv.appendChild(child);

    this.replaceChildren(leftDiv, centerDiv, rightDiv);
  }
}

export class YrNavbarStd extends HTMLElement {
  connectedCallback() {
    this.setAttribute('role', 'navigation');
    this.restructure();
  }

  private restructure() {
    if (this.querySelector('.yr-navbar-left')) return;

    const leftDiv = document.createElement('div');
    leftDiv.className = 'yr-navbar-left';

    const logo = document.createElement('a');
    logo.href = '/';
    logo.className = 'yr-brand';
    const img = document.createElement('img');
    img.src = this.getAttribute('logo') || '/logo.png';
    img.alt = this.getAttribute('brand') || 'Yardrail';
    logo.appendChild(img);
    leftDiv.appendChild(logo);

    const centerDiv = document.createElement('div');
    centerDiv.className = 'yr-navbar-center';
    const center = this.querySelector('[slot="center"]');
    if (center) centerDiv.appendChild(center);

    const rightDiv = document.createElement('div');
    rightDiv.className = 'yr-navbar-right';
    const remaining = Array.from(this.childNodes).filter(
      (n) => n !== leftDiv && n !== centerDiv,
    );
    for (const child of remaining) rightDiv.appendChild(child);

    if (!this.hasAttribute('authed')) {
      const loginHref = this.getAttribute('login-href') || '/login';
      const signupHref = this.getAttribute('signup-href') || '/signup';

      const login = document.createElement('a');
      login.href = loginHref;
      login.className = 'yr-navbar-login';
      login.textContent = 'Log in';
      rightDiv.appendChild(login);

      const signup = document.createElement('a');
      signup.href = signupHref;
      signup.className = 'yr-navbar-signup';
      signup.textContent = 'Sign up';
      rightDiv.appendChild(signup);
    }

    this.replaceChildren(leftDiv, centerDiv, rightDiv);
  }
}

customElements.define('yr-navbar', YrNavbar);
customElements.define('yr-navbar-std', YrNavbarStd);
