import React from 'react';
import Layout from '@theme/Layout';
import { WriteInReact } from "../../components/WriteInReact";
import styles from "./landing.module.css";

const NewLanding: React.FC = () => {
  return (
    <Layout>
      <div className={styles.content}>
        <WriteInReact />
      </div>
    </Layout>
  );
}

export default NewLanding;