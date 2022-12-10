// importがあるのでモジュール"拡張"を書くことができる (importを消すならexport {};必須)
import type { SuperTest, Test } from 'supertest';

declare global {
  // environ.mjs内でセット(setup.js内で作成)
  var request: SuperTest<Test>;
}