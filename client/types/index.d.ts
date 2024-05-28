declare global {
  interface Provider {
    id: number;
    firstname: string;
    lastname: string;
    suffix: string;
    bio: string;
    email: string;
    phone: string;
    image: string;
    languages: string[];
    services: string[];
  }

  type Allergy = {
    description: string;
    name: string;
  };

  type Prescription = {
    name: string;
    dosage: string;
    start: Date;
    end: Date;
    frequency: string;
    providerId: string;
  };

  interface Patient {
    id: number;
    firstname: string;
    lastname: string;
    email: string;
    language: string;
    birth: Date;
    gender: 'male' | 'female';
    allergies: Allergy[];
    prescriptions: Prescription[];
  }
}

export {};
