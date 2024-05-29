import bcrypt from "bcryptjs";
const fixedSalt = "$2a$10$1234567890123456789012";
interface LoginResponse {
  id: string;
  token: string;
}

const pid = useCookie("pid");
const token = useCookie("token");

const hashPassword = async (password: string) => {
  return await bcrypt.hash(password, fixedSalt);
}

export function useAuth() {
  const isLoggedIn = ref(false);

  async function registerPatient(patient: Patient): Promise<boolean> {
    patient.password = await hashPassword(patient.password || '');
    return $fetch("/api/patient", {
      method: "POST",
      body: JSON.stringify(patient),
    }).then((res) => {
      isLoggedIn.value = true;
      return true;
    }).catch((err) => {
      console.log(err);
      isLoggedIn.value = false;
      return false;
    });
  }

  async function registerProvider(provider: Provider): Promise<boolean> {
    provider.password = await hashPassword(provider.password || '');
    return $fetch("/api/provider", {
      method: "POST",
      body: JSON.stringify(provider),
    }).then((res) => {
      isLoggedIn.value = true;
      return true;
    }).catch((err) => {
      console.log(err);
      isLoggedIn.value = false;
      return false;
    });
  }

  async function loginPatient(patient: Login): Promise<boolean | undefined> {
    return $fetch("/api/patient/login", {
      method: "POST",
      body: JSON.stringify(patient),
    }).then((res) => {
      const loginResponse: LoginResponse = res as LoginResponse;
      pid.value = loginResponse.id;
      token.value = loginResponse.token;
      isLoggedIn.value = true;
      return true;
    }).catch((err) => {
      console.log(err);
      isLoggedIn.value = false;
      return false;
    });
  }

  async function loginProvider(provider: Login): Promise<boolean | undefined> {
    return $fetch("/api/provider/login", {
      method: "POST",
      body: JSON.stringify(provider),
    }).then((res) => {
      const loginResponse: LoginResponse = res as LoginResponse;
      pid.value = loginResponse.id;
      token.value = loginResponse.token;
      isLoggedIn.value = true;
      return true;
    }).catch((err) => {
      console.log(err);
      isLoggedIn.value = false;
      return false;
    });
  }

  async function getPatientData(): Promise<Patient | null> {
    return $fetch(`/api/patient/${pid.value}`, {
      headers: {
        Authorization: `Bearer ${token.value}`,
      },
    }).then((res) => {
      return res as Patient;
    }).catch((err) => {
      console.log(err);
      return null;
    })
  }

  async function getProviderData(): Promise<Provider | null> {
    return $fetch(`/api/provider/${pid.value}`, {
      headers: {
        Authorization: `Bearer ${token.value}`,
      },
    }).then((res) => {
      return res as Provider;
    }).catch((err) => {
      console.log(err);
      return null;
    })
  }

  async function logoutUser() {
    isLoggedIn.value = false;
    pid.value = "";
    token.value = "";
  }

  return { registerPatient, registerProvider, loginPatient, loginProvider, getPatientData, getProviderData, logoutUser };
}