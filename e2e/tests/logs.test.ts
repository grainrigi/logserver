import { cleanupDB, extractId } from '../util/common';
import { createContest, deleteContest } from '../util/contests';

describe('logs', () => {
  beforeAll(async () => {
    await cleanupDB();
  });

  let cid: number;

  const log = {
    time: new Date().toISOString(),
    call: 'JJ1SHQ',
    rst: '599',
    rcvd: '100109M',
    band: '7',
    mode: 'CW',
    pwr: 'M',
    note: 'test',
  };
  const log2 = {
    time: new Date().toISOString(),
    call: 'JI1SDI',
    rst: '599',
    rcvd: '100109M',
    band: '14',
    mode: 'CW',
    pwr: 'M',
    note: 'test2',
  };
  let id: number;
  let id2: number;

  it('can be created', async () => {
    cid = await createContest();
    await request
      .post(`/contests/${cid}/logs`)
      .send(log)
      .expect(201)
      .then((res) => {
        id = extractId(res);
      });
    await request
      .post(`/contests/${cid}/logs`)
      .send(log2)
      .expect(201)
      .then((res) => {
        id2 = extractId(res);
      });
  });

  it('can be read', async () => {
    await request
      .get(`/contests/${cid}/logs`)
      .expect(200)
      .then((res) => {
        expect(res.body).toContainEqual(expect.objectContaining(log));
        expect(res.body).toContainEqual(expect.objectContaining(log2));
      });
  });

  it('can be updated', async () => {
    log.call = 'JJ1UZH';
    await request
      .put(`/contests/${cid}/logs/${id}`)
      .send(log)
      .expect(204);
    await request
      .get(`/contests/${cid}/logs`)
      .expect(200)
      .then((res) => {
        expect(res.body).toContainEqual(expect.objectContaining(log));
        expect(res.body).toContainEqual(expect.objectContaining(log2));
      });
  });

  it('can be deleted', async () => {
    await request
      .delete(`/contests/${cid}/logs/${id}`)
      .expect(204);
    await request
      .get(`/contests/${cid}/logs`)
      .expect(200)
      .then((res) => {
        expect(res.body).not.toContainEqual(expect.objectContaining(log));
        expect(res.body).toContainEqual(expect.objectContaining(log2));
      });
  });

  it('can be cascade deleted', async () => {
    await deleteContest(cid);
  });
});