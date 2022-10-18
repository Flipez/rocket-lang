import React from "react";
import Tabs from "@theme/Tabs";
import TabItem from "@theme/TabItem";
import CodeBlock from "@theme/CodeBlock";
import styles from "./Highlights.module.css";
import Link from "@docusaurus/Link";
import clsx from "clsx";
import { GetStarted } from "./GetStarted";

const WelcomeCode = `
🚀 > puts("hello from rocket-lang!")
"hello from rocket-lang!"
=> nil

🚀 > langs = ["ruby", "go", "crystal", "python", "php"]
=> ["ruby", "go", "crystal", "python", "php"]

🚀 > langs.yeet()
=> "php"

🚀 > langs.yoink("rocket-lang")
=> nil

🚀 > langs
=> ["ruby", "go", "crystal", "python", "rocket-lang"]
`

function Welcome() {
  return (
    <section className={clsx(styles.section)}>
      <div className="container">
        <div className="row">
          <div className="col col--6">
            <h1 className={styles.writeincsstitle}>
              It's not <br /> rocket science.
            </h1>
            <p>
              Use some of the syntax features of Ruby (but worse) and create programs that will maybe perform better.
            </p>

            < GetStarted />
          </div>
          <div className="col col--6">
            <CodeBlock language="js" children={WelcomeCode} />
          </div>
        </div>
      </div>
    </section>
  );
}




const JSONExample = `
🚀 > JSON.parse('{"test": 123}')
=> {"test": 123.0}

🚀 > a = {"test": 1234}
=> {"test": 1234}

🚀 > a.to_json()
=> '{"test":1234}'
`;

const HTTPExample = `
def test()
  puts(request["body"])
  return("test")
end

HTTP.handle("/", test)

HTTP.listen(3000)
`;

const ClosuresExample = `
newGreeter = def (greeting)
  return def (name)
           puts(greeting + " " + name)
         end
end

hello = newGreeter("Hello");
hello("dear, future Reader!");
`;

const BuiltinList = [
  {
    label: "JSON",
    value: "json",
    content: <CodeBlock language="js" children={JSONExample} />,
  },
  {
    label: "HTTP",
    value: "http",
    content: <CodeBlock language="js" children={HTTPExample} />,
  },
  {
    label: "Closures",
    value: "closures",
    content: <CodeBlock language="js" children={ClosuresExample} />,
  },
];

function Builtins() {
  return (
    <section className={styles.section}>
      <div className="container">
        <div className="row">
          <div className="col col--6">
            <Tabs defaultValue="json" values={BuiltinList}>
              {BuiltinList.map((props, idx) => {
                return (
                  <TabItem key={idx} value={props.value}>
                    {props.content}
                  </TabItem>
                );
              })}
            </Tabs>
          </div>
          <div className="col col--6">
            <h2>
              Strong and stable <span className="highlight">builtins</span>
            </h2>
            <p>
              RocketLang ships some neat builtins such as handling HTTP requests and responses,
              marshalling and unmashalling of JSON objects.
            </p>
            <p>
              It also comes with support of closures, modules and context sensitive variables in order
              to create complex programs.
            </p>
          </div>
        </div>
      </div>
    </section>
  );
}

const ObjectExample = `
🚀 > "test".type()
=> "STRING"

🚀 > true.wat()
=> BOOLEAN supports the following methods:
                plz_s()

🚀 > 1.methods()
=> ["plz_s", "plz_i", "plz_f"]
`;

function EverythingObject() {
  return (
    <section className={clsx(styles.section)}>
      <div className="container">
        <div className="row">
          <div className="col">
            <h2>
              <span className="highlight">Everything</span> is an object
            </h2>
          </div>
        </div>
        <div className="row">
          <div className="col col--6">
            <p className={styles.paddingTopLg}>
              Inspired by Ruby, in RocketLang everything is an object.
            </p>
            <p>
              This allows to threat unknown input in the same way and figuring out what kind of
              information your function passes on the go.
              Every object supports the same minimum default subset of methods to achive this.
            </p>
          </div>
          <div className="col col--6">
            <CodeBlock language="js" children={ObjectExample} />
          </div>
        </div>
      </div>
    </section>
  );
}

export default function Highlights() {
  return (
    <span>
      {Welcome()}
      {Builtins()}
      {EverythingObject()}
    </span>
  );
}
