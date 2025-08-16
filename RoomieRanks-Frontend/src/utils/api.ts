const API_BASE_URL = 'http://localhost:8080/api/v1';

export interface User {
  username: string;
  name: string;
  email: string;
  household_id: number;
  title: string;
  level: number;
  xp: number;
}

export interface ApiError {
  error: string;
}

export async function fetchUser(username: string): Promise<User> {
  const url = `${API_BASE_URL}/user/${username}`;
  console.log('🔍 Fetching user data from:', url);
  console.log('🔍 Username being requested:', username);
  
  try {
    const response = await fetch(url, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
      },
    });

    console.log('📡 Response status:', response.status);
    console.log('📡 Response status text:', response.statusText);
    console.log('📡 Response headers:', Object.fromEntries(response.headers.entries()));

    if (!response.ok) {
      console.error('❌ Response not OK. Status:', response.status);
      
      let errorMessage = 'Failed to fetch user data';
      try {
        const errorData: ApiError = await response.json();
        console.error('❌ Error data from server:', errorData);
        errorMessage = errorData.error || errorMessage;
      } catch (parseError) {
        console.error('❌ Could not parse error response:', parseError);
        const textError = await response.text();
        console.error('❌ Raw error response:', textError);
      }
      
      throw new Error(errorMessage);
    }

    const userData = await response.json();
    console.log('✅ Successfully fetched user data:', userData);
    return userData;
  } catch (error) {
    console.error('💥 Network or other error:', error);
    throw error;
  }
}

export async function updateUserTitle(username: string, title: string): Promise<void> {
  const response = await fetch(`${API_BASE_URL}/user/${username}/title`, {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ title }),
  });

  if (!response.ok) {
    const errorData: ApiError = await response.json();
    throw new Error(errorData.error || 'Failed to update title');
  }
}

export async function addUserXP(username: string, xp: number): Promise<void> {
  const response = await fetch(`${API_BASE_URL}/user/${username}/xp`, {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ xp }),
  });

  if (!response.ok) {
    const errorData: ApiError = await response.json();
    throw new Error(errorData.error || 'Failed to add XP');
  }
}
