import React from 'react';
import CodeBlock from '@theme/CodeBlock';

class CodeExample extends React.Component {
  constructor(props){
    super(props);
    this.state = {
      output: '',
    }
  }

  render() {

    if (window && document) {
      const body = document.getElementsByTagName('body')[0]

      const wasm_script = document.createElement('script')
      wasm_script.src = 'https://play.rocket-lang.org/wasm_exec.js'
      body.appendChild(wasm_script)
      
      wasm_script.addEventListener('load', () => {

        let outputBuf = '';
		    const decoder = new TextDecoder("utf-8");
		    globalThis.fs.writeSync = (fd, buf) => {
		    	outputBuf += decoder.decode(buf);
		    	const nl = outputBuf.lastIndexOf("\n");
		    	if (nl != -1) {
		    		this.state.output += outputBuf.substr(0, nl + 1);
		    		window.scrollTo(0, document.body.scrollHeight);
		    		outputBuf = outputBuf.substr(nl + 1);
		    	}
		    	return buf.length;
		    };

        fetch('https://play.rocket-lang.org/main.wasm').then(response => response.arrayBuffer()).then((bin) => {
          const go = new Go();
	        go.argv = ['rocket-lang', '-e', this.props.code];
	        go.exit = (code) => {
		        if (code > 0)
		      	  this.state.output += 'Exit ' + code + '\n';
		       };

          WebAssembly.instantiate(bin, go.importObject).then((result) => {
				  	go.run(result.instance);
				  });
        })
      })
    }

    return (
      <div
      style={{
        backgroundColor: '#303846',
        borderRadius: 5,
      }}>
        <CodeBlock
          language="js"
          showLineNumbers>
          {this.props.code}
        </CodeBlock>
        <CodeBlock language="plain" title='Output'>
          {this.state.output}
        </CodeBlock>
      </div>
    );
  }
}

export default CodeExample