import React from 'react';
import CodeBlock from '@theme/CodeBlock';

class CodeBlockSimple extends React.Component {
  constructor(props){
    super(props);
  }

  render() {
    return (
      <div
      style={{
        backgroundColor: '#303846',
        borderRadius: 5,
      }}>
        <CodeBlock language="js" showLineNumbers>
          {this.props.input}
        </CodeBlock>
        { this.props.output ? (
        <CodeBlock language="js" title='Output'>
          {this.props.output}
        </CodeBlock>) : ("") }
      </div>
    );
  }
}

export default CodeBlockSimple