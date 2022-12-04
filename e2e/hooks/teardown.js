module.exports = async function (globalConfig, projectConfig) {
  await global.serverProcess.kill('SIGTERM');
}