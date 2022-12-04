import { TestEnvironment as NodeEnvironment } from 'jest-environment-node';
import chalk from 'chalk';

class GoServerEnvironment extends NodeEnvironment {
  logs = "";
  onData = (chunk) => this.logs += chunk;

  async setup() {
    await super.setup();
    // requestをテスト内で使えるようにする
    this.global.request = global.request;
  }

  async handleTestEvent(event, state) {
    switch (event.name) {
      case 'test_start':
        this.logs = '';
        serverProcess.stdout?.addListener('data', this.onData);
        serverProcess.stderr?.addListener('data', this.onData);
        break;
      case 'test_done':
        // 失敗時にサーバーのログを出力(当該リクエストに関係あるところだけ)
        if (event.test.errors?.length > 0) {
          console.log(chalk.bold('Server Output For', `"${event.test.parent.name} › ${event.test.name}"`));
          console.log(this.logs);
        }
        serverProcess.stdout?.removeListener('data', this.onData);
        serverProcess.stderr?.removeListener('data', this.onData);
        break;
    }
  }
}

export default GoServerEnvironment;
