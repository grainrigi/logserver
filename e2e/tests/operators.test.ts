import { Test, Response } from 'supertest';
import { v4 as uuidv4 } from 'uuid';

describe('operators', () => {
  const opname = uuidv4();
  const op = { name: opname, license: 1 };
  let opid = 0;

  const has = (yes: boolean) => (res: Response) =>
      (yes ? expect(res.body) : expect(res.body).not).toContainEqual(expect.objectContaining(op));

  it('can be created', () => {
    return request
      .post('/operators')
      .send(op)
      .expect(201);
  });

  it('cannot be created with invalid value', () => {
    return request
      .post('/operators')
      .send({})
      .expect(400);
  });

  it('can be read', () => {
    return request
      .get('/operators')
      .expect(200)
      .then((res) => {
        has(true)(res);
        opid = res.body.find((v: any) => v.name === opname).id;
      });
  });

  it('can be updated', async () => {
    op.name = uuidv4();
    op.license = 2;
    await request
      .put('/operators/' + opid)
      .send({ ...op, id: opid })
      .expect(204);
    await request
      .get('/operators')
      .expect(200)
      .then(has(true));
  })

  it('can be deleted', async () => {
    await request
      .delete('/operators/' + opid)
      .expect(204);
    // 削除されたか確認
    await request
      .get('/operators')
      .expect(200)
      .then(has(false));
  });
});