type HttpMethod = 'GET' | 'POST' | 'PUT' | 'DELETE' | 'PATCH' | 'HEAD' | 'OPTIONS';

interface FetchOptions {
  url: string;
  method?: HttpMethod;
  body?: Record<string, unknown> | null;
  headers?: Record<string, string>;
}

// 定义后端统一响应格式
interface ApiResponse<T = unknown> {
  code: number;
  msg: string;
  data: T;
}

async function fetchRequest<T = unknown>(options: FetchOptions): Promise<T> {
  const { 
    url,
    method = 'GET',
    body = null,
    headers = {}
  } = options;

  const mergedHeaders = {
    'Content-Type': 'application/json',
    ...headers
  };

  const config: RequestInit = {
    method,
    headers: new Headers(mergedHeaders),
    credentials: 'same-origin'
  };

  if (method !== 'GET' && method !== 'HEAD' && body) {
    config.body = JSON.stringify(body);
  }

  try {
    const response = await fetch(url, config);
    
    // 强制解析为JSON格式（根据你的业务需求）
    const result = await response.json();

    // 业务状态码判断
    if (result.code !== 0) {
      throw new Error(result.msg);
    }

    // 返回实际业务数据
    return result;
  } catch (error) {
    console.error('请求处理失败:', error);
    throw error instanceof Error ? error : new Error('未知错误');
  }
}

export default fetchRequest