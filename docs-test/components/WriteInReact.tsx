import React from "react";
import { GetStarted } from "./GetStarted";
import styles from "./writeinreact.module.css";

export const WriteInReact: React.FC = () => {
  return (
    <div className={styles.writeincss}>
      <div style={{ flex: 1 }}>
        <h1 className={styles.writeincsstitle}>
          It's not <br /> rocket science.
        </h1>
        <p>
          Use some of the syntax features of Ruby (but worse) and create programs that will maybe perform better.
        </p>

        < GetStarted />
      </div>
      <div className={styles.writeright}>
        <img src="/img/landing_page_terminal.svg"/>
      </div>
    </div>
  );
};
