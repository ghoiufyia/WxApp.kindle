const app = getApp();

Page({
  data: {},
  onLoad() {
    app.getUserInfo().then(
      user => {
        this.setData({
          user,
        });
      },
      () => {
        // 获取用户信息失败
      }
    );
  },
});
