interface TokenResult {
  name: string;
  accessToken: string;
  refreshToken: string;
}

interface CaptchaResult {
  captchaId: string;
  image: string;
  tile: string;
  y: number;
  width: number;
}

interface SignIn {
  name?: string;
  password?: string;
  captchaId: string;
  captchaX: number;
  captchaY: number;
}
