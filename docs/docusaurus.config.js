// @ts-check
// Note: type annotations allow type checking and IDEs autocompletion

//const lightCodeTheme = require('prism-react-renderer/themes/github');
//const darkCodeTheme = require('prism-react-renderer/themes/dracula');
const codeTheme = require('./components/CodeBlockTheme');

/** @type {import('@docusaurus/types').Config} */
const config = {
  title: '🚀🇱🅰🆖',
  tagline: '',
  url: 'https://rocket-lang.org',
  baseUrl: '/',
  onBrokenLinks: 'warn',
  onBrokenMarkdownLinks: 'warn',
  favicon: 'img/favicon.ico',
  // Even if you don't use internalization, you can use this field to set useful
  // metadata like html lang. For example, if your site is Chinese, you may want
  // to replace "en" with "zh-Hans".
  i18n: {
    defaultLocale: 'en',
    locales: ['en'],
  },

  future: {
    experimental_faster: true,
    v4: {
      removeLegacyPostBuildHeadAttribute: true,
    }
  },

  presets: [
    [
      'classic',
      /** @type {import('@docusaurus/preset-classic').Options} */
      ({
        docs: {
          sidebarCollapsible: false,
          sidebarPath: require.resolve('./sidebars.js'),
          editUrl:
            'https://github.com/flipez/rocket-lang/tree/main/docs/',
            lastVersion: 'v0.22.1',
        },
        theme: {
          customCss: require.resolve('./src/css/custom.css'),
        },
      }),
    ],
  ],

  themeConfig:
    /** @type {import('@docusaurus/preset-classic').ThemeConfig} */
    ({
      algolia: {
        appId: 'P0XYEH7IPQ',
        apiKey: 'd5f8e27e198519529222235cec0febfe',
        indexName: 'rocket-lang',
      },
      navbar: {
        title: 'RocketLang',
        logo: {
          alt: 'RocketLang',
          src: 'img/logo.png',
        },
        items: [
          {
            type: 'docsVersionDropdown',
            position: 'left',
            dropdownActiveClassDisabled: true,
          },
          {
            to: 'docs',
            position: 'left',
            label: 'Documentation',
          },
          {to: '/blog', label: 'Blog', position: 'left'},
          {to: 'https://play.rocket-lang.org', label: 'Playground', position: 'left'},
          {
            type: 'search',
            position: 'right',
          },
          {
            href: 'https://github.com/flipez/rocket-lang',
            label: 'GitHub',
            position: 'right',
          },
        ],
      },
      footer: {
        style: 'dark',
        links: [],
        copyright: `Copyright © ${new Date().getFullYear()} RocketLang.org`,
      },
      prism: {
        theme: codeTheme,
        darkTheme: codeTheme,
      },
    }),
};

module.exports = config;
