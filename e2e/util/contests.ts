import { extractId } from './common';

export async function createContest(c?: object): Promise<number> {
  const contest = {
    name: '全市全郡コンテスト',
    startTime: new Date().toISOString(),
    endTime: new Date(new Date().valueOf() + 100000).toISOString(),
    type: 1,
    cfg: "power -HH-H-H-HHMMLPPP",
    call: "JA1YAD",
    ...c,
  }

  return await request
    .post('/contests')
    .send(contest)
    .expect(201)
    .then((res) => {
      return extractId(res);
    });
}

export async function deleteContest(id: number) {
  await request
    .delete('/contests/' + id)
    .expect(204);
}