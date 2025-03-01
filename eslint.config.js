import pluginJs from "@eslint/js";
import globals from "globals";

export default [
  pluginJs.configs.recommended,
  {
    languageOptions: {
      globals: {
        ...globals.node,
      },
    },
    files: ["**/*.ts"],
    rules: {
      "no-unused-vars": "warn",
      "no-undef": "warn",
      quotes: [2, "single", { avoidEscape: true }],
    },
  },
];
