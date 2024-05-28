declare global {
  interface Provider {
    id: string;
    firstname: string;
    lastname: string;
    suffix: string;
    bio: string;
    email: string;
    phone: string;
    image: string;
    languages: string[];
    services: string[];
    password: ?string;
  }

  interface Allergy {
    description: string;
    name: string;
  }

  interface Prescription {
    name: string;
    dosage: string;
    start: Date;
    end: Date;
    frequency: string;
    providerId: string;
  }

  interface Patient {
    id: string;
    firstname: string;
    lastname: string;
    email: string;
    language: string;
    birth: Date;
    gender: 'male' | 'female';
    allergies: Allergy[];
    prescriptions: Prescription[];
    password: ?string;
  }

  interface Appointment {
    id: string;
    date: Date;
    startTime: string;
    endTime: string;
    provider: string;
    patient: string;
    service: number;
    description: string;
  }

  interface Login {
    email: string;
    password: string;
  }
}

export {};
