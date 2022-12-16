import { cleanupDB } from '../util/common';
import { createContest, deleteContest } from '../util/contests';

describe('contests', () => {
  beforeAll(async () => {
    await cleanupDB();
  });

  const contest = {
    name: '全市全郡コンテスト',
    startTime: new Date().toISOString(),
    endTime: new Date(new Date().valueOf() + 100000).toISOString(),
    type: 1,
    cfg: "power -HH-H-H-HHMMLPPP",
    call: "JA1YAD",
  };

  let id: number;

  it('can be created', async () => {
    id = await createContest(contest);
  });

  it('cannot be created with invalid value', async () => {
    // typeは1か2
    await request
      .post('/contests')
      .send({ ...contest, type: 3 })
      .expect(400);
  });

  it('can be read', () => {
    return request
      .get('/contests')
      .expect(200)
      .then((res) => {
        expect(res.body).toContainEqual({ ...contest, id });
      });
  });

  it('can be solely read', () => {
    return request
      .get('/contests/' + id)
      .then((res) => {
        expect(res.body).toEqual({ ...contest, id });
      });
  });

  it('can be updated', async () => {
    contest.name = "ぜんしぜんぐん"
    await request
      .put('/contests/' + id)
      .send(contest)
      .expect(204);
    // 正しく更新されたか確認
    await request
      .get('/contests')
      .expect(200)
      .then((res) => {
        expect(res.body).toContainEqual({ ...contest, id });
      });
  });

  it('callsign is coerced to upper case', async () => {
    contest.call = 'JA1ZGP';
    await request
      .put('/contests/' + id)
      .send({ ...contest, call: 'ja1zgp' })
      .expect(204);
    await request
      .get('/contests')
      .expect(200)
      .then((res) => {
        expect(res.body).toContainEqual({ ...contest, id });
      });
  });

  it('can be deleted', async () => {
    await deleteContest(id);
    await request
      .get('/contests')
      .expect(200)
      .then((res) => {
        expect(res.body).not.toContainEqual(expect.objectContaining({ id: 1 }));
      });
  });
});
