// importがあるのでモジュール"拡張"を書くことができる (importを消すならexport {};必須)
import type { SuperTest, Test } from 'supertest';

declare global {
  var request: SuperTest<Test>;
}