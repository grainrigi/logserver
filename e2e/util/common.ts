import supertest from 'supertest';
import { execa } from '@esm2cjs/execa';

export function extractId(res: supertest.Response): number {
  expect(res.body).toEqual(expect.objectContaining({ id: expect.any(Number) }));
  return res.body.id;
}

export async function cleanupDB() {
  // migrate down & up
  const migrateOptions = ['-path', '../migrations', '-database', process.env.DB_URL!];
  await execa('/usr/bin/migrate', [...migrateOptions, 'down', '-all']);
  await execa('/usr/bin/migrate', [...migrateOptions, 'up']);
}