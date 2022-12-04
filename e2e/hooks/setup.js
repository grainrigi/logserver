const { setTimeout } = require('timers/promises');
const supertest = require('supertest');

module.exports = async function (globalConfig) {
  const { execa } = await import('execa');
  const chalk = (await import('chalk')).default;
  global.serverProcess = execa('../logserver', [], {
    env: { ...process.env, BUNDEBUG: 2 },
  });

  // wait until the server come up
  const ac = new AbortController();

  let logs = "";
  const comeupPromise = new Promise((resolve) => {
    const hook = () => {
      if (logs.includes('http server started on')) {
        resolve('started');
      }
    };
    global.serverProcess.stdout.on('data', chunk => { logs += chunk; hook(); }, { signal: ac });
    global.serverProcess.stderr.on('data', chunk => { logs += chunk; hook(); }, { signal: ac });
  });

  const result = await Promise.race([
    comeupPromise,
    setTimeout(5000, 'timedout', { signal: ac.signal }),
    global.serverProcess.catch((reason) => reason),
  ]);

  // 正常起動
  if (result === 'started') {
    // サーバーが立ち上がったので、テストの準備をする
    ac.abort();
    global.request = supertest('http://localhost:3030');
    return;
  }

  // 異常系処理
  console.log(); // line break
  if (result === 'timedout') {
    console.log(chalk.red.bold("Server didn't come up after 5000ms. Aborting"))
    await global.serverProcess.kill('SIGTERM', {
      forceKillAfterTimeout: 5000,
    });
    console.log(chalk.red.bold("Output:"));
    console.log(chalk.red.bold(logs));
  } else if (typeof result === 'object' && typeof result.exitCode === 'number') {
    console.log(chalk.red.bold("Server exited with exitCode " + result.exitCode + ".\nOutput:"));
    console.log(chalk.red.bold(logs));
  } else if (result !== 'started') {
    console.log(chalk.red.bold("Unknown result:"), result);
  }
  process.exit(1);
}