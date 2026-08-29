import React from "react";
import Tabs from "@theme/Tabs";
import TabItem from "@theme/TabItem";
import CodeBlock from "@theme/CodeBlock";
import styles from "./Highlights.module.css";
import Link from "@docusaurus/Link";
import clsx from "clsx";
import { GetStarted } from "./GetStarted";

const WelcomeCode = `🚀 » puts("hello from rocket-lang!")
hello from rocket-lang!
» nil

🚀 » langs = ["ruby", "go", "crystal", "python", "php"]
» ["ruby", "go", "crystal", "python", "php"]

🚀 » langs.pop()
» "php"

🚀 » langs.push("rocket-lang")
» ["ruby", "go", "crystal", "python", "rocket-lang"]

🚀 » langs.sort().first()
» "crystal"
`

function Welcome() {
  return (
    <section className={clsx(styles.section)}>
      <div className="container">
        <div className="row">
          <div className="col col--6">
            <h1 className={styles.writeincsstitle}>
              It's not <br /> rocket science. Is it?
            </h1>
            <p>
              Started as a joke to learn the inner mechanics of a interpreted language,
              <span className="highlight2"> RocketLang </span>
              quickly evolved into a long running project with more and more ambitious
              goals, such as being used for&nbsp;
              <a href="https://adventofcode.com">Advent of Code</a> to test it's limits.
            </p>
          </div>
          <div className="col col--6">
            <CodeBlock language="js" children={WelcomeCode} />
          </div>
        </div>
      </div>
    </section>
  );
}




const JSONExample = `🚀 » JSON.parse('{"test": 123}')
» {"test": 123.0}


🚀 » a = {"test": 1234}
» {"test": 1234}


🚀 » a.to_json()
» "{\\"test\\":1234}"
`;

const HTTPExample = `def hello()
  puts(request["body"])
  response["body"] = "hello " + request["method"]
end


server = HTTP.new()
server.handle("/", hello)


server.listen(3000)
`;

const MathExample = `🚀 » Math.E
» 2.718281828459045


🚀 » Math.Pi
» 3.141592653589793


🚀 » Math.sqrt(3.0 * 3.0 + 4.0 * 4.0)
» 5.0
`;

const TimeExample = `🚀 » Time.format(Time.unix(), "Mon Jan _2 15:04:05 2006")
» "Sat Aug 29 22:09:48 2026"
🚀 » Time.format(Time.unix(), "%a %b %e %H:%M:%S %Y")
» "Sat Aug 29 22:09:48 2026"


🚀 » Time.parse("2022-03-23", "2006-01-02")
» 1647993600
🚀 » Time.parse("2022-03-23", "%Y-%m-%d")
» 1647993600
`;

const ClosuresExample = `newGreeter = def (greeting)
  return def (name)
    puts(greeting + " " + name)
  end
end

hello = newGreeter("Hello")
hello("dear, future Reader!")

Hello dear, future Reader!
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
    label: "Math",
    value: "math",
    content: <CodeBlock language="js" children={MathExample} />,
  },
  {
    label: "Time",
    value: "time",
    content: <CodeBlock language="js" children={TimeExample} />,
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
              Strong and stable <span className="highlight2">builtins</span>
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

const ObjectExample = `🚀 » "test".type()
» "STRING"

🚀 » true.to_s()
» "true"

🚀 » 1.4.to_s()
» "1.4"
`;

function EverythingObject() {
  return (
    <section className={clsx(styles.section)}>
      <div className="container">
        <div className="row">
          <div className="col col--6">
            <h2>
              <span className="highlight2">Everything</span> is an object
            </h2>
          </div>
        </div>
        <div className="row">
          <div className="col col--6">
            <p className={styles.paddingTopLg}>
              Inspired by Ruby, in RocketLang everything is an object.
            </p>
            <p>
              This allows to treat unknown input in the same way and figuring out what kind of
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

function Quickstart() {
  return (
    <section className={clsx(styles.section)}>
      <div className="container">
        <div className="row">
          <div className="col">
            <h2>
              <span className="highlight2">Get started</span> quickly
            </h2>
          </div>
        </div>
        <div className="row">
          <div className="col">
            <div className={styles.centerjustify}>
              <GetStarted />
            </div>
            <div className={styles.centerjustify}>
              <p>
                You can also install RocketLang via APT, RPM, as a binary or from source.
                See <a href="/docs/install">&nbsp;the install guide&nbsp;</a> for more detailed information.
              </p>
            </div>
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
      {Quickstart()}
      {Builtins()}
      {EverythingObject()}
    </span>
  );
}
