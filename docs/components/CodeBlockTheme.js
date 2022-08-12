var theme = {
  plain: {
    color: "rgba(255,90,121,255)",
    backgroundColor: "rgba(48,56,70,255)",
  },
  styles: [
    {
      types: ["prolog", "constant", "builtin", "boolean"],
      style: {
        color: "rgb(189, 147, 249)",
      },
    },
    {
      types: ["inserted", "function"],
      style: {
        color: "rgba(255,157,39,255)",
      },
    },
    {
      types: ["deleted"],
      style: {
        color: "rgb(255, 85, 85)",
      },
    },
    {
      types: ["changed"],
      style: {
        color: "rgb(255, 184, 108)",
      },
    },
    {
      types: ["punctuation", "symbol"],
      style: {
        color: "rgb(248, 248, 242)",
      },
    },
    {
      types: ["string", "char", "tag", "selector"],
      style: {
        color: "rgba(79,209,217,255)",
      },
    },
    {
      types: ["keyword", "variable"],
      style: {
        color: "rgba(253,245,22,255)",
        fontStyle: "italic",
      },
    },
    {
      types: ["operator"],
      style: {
        color: "rgba(253,245,22,255)",
      },
    },
    {
      types: ["comment"],
      style: {
        color: "rgb(98, 114, 164)",
      },
    },
    {
      types: ["attr-name"],
      style: {
        color: "rgba(253,245,22,255)",
      },
    },
  ],
};

module.exports = theme;
