import React, { useState } from "react";
import styles from "./get-started.module.css";

export const GetStarted: React.FC = () => {
  const [clicked, setClicked] = useState<number | null>(null);
  return (
    <>
      <div className={styles.myrow}>
        <div style={{ position: "relative" }}>
          {clicked ? (
            <div key={clicked} className={styles.copied}>
              Copied!
            </div>
          ) : null}
          <div
            className={styles.codeblock}
            onClick={() => {
              navigator.clipboard.writeText("brew install flipez/homebrew-tap/rocket-lang");

              setClicked(Date.now());
            }}
            title="Click to copy"
          >
            $ brew install flipez/homebrew-tap/rocket-lang
          </div>
        </div>
      </div>
      <br />
      <br />
      <br />
      <br />
    </>
  );
};
